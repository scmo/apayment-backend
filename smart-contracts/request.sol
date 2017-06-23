pragma solidity ^0.4.11;


import "rbac.sol";


contract mortal {
  /* Define variable owner of the type address*/
  address owner;

  uint public created;

  uint public  modified;
  /* this function is executed at initialization and sets the owner of the contract */
  function mortal() {owner = msg.sender;}

  /* Function to recover the funds on the contract */
  function kill() {if (msg.sender == owner) selfdestruct(owner);}

  function setCreated(){
    created = block.timestamp;
  }

  function setModified(){
    modified = block.timestamp;
  }
}


contract Request is mortal {
  address public inspectorAddress;

  uint16[] public contributionCodes;

  string public remark;

  RBAC rbac;

  struct Lack {
  uint16 contributionCode;
  string controlCategoryId;
  uint16 pointGroupCode;
  string controlPointId;
  int64 lackId;
  uint8 points;
  }

  struct LackNew {
  uint16 contributionCode;
  int64 controlCategoryId;
  uint16 pointGroupCode;
  int64 controlPointId;
  int64 lackId;
  uint8 points;
  }

  uint public numLacks;

  mapping (uint => Lack) public lacks;

  mapping (uint => LackNew) public newLacks;

  function Request(uint16[] _contributionCodes, string _remark) public {
    rbac = RBAC("0x0a3C03F7dE1CeFc9284358503225a0F269a4d868");
    contributionCodes = _contributionCodes;
    remark = _remark;
    //setGVE(gves[0], gves[1], gves[2], gves[3], gves[4], gves[5], gves[6], gves[7], gves[8]);
    setCreated();
  }

  function setInspectorId(address _inspectorAddress){
    require(rbac.isAdmin(msg.sender) || rbac.isCantonEmployee(msg.sender));
    require(rbac.isInspector(_inspectorAddress));
    inspectorAddress = _inspectorAddress;
    setModified();
  }

  function addLack(uint16 _contributionCode, string _controlCategoryId, uint16 _pointGroupCode, string _controlPointId, int64 _lackId, uint8 _points) {
    require(msg.sender == inspectorAddress);
    uint lacksIndex = numLacks++;
    lacks[lacksIndex] = Lack(_contributionCode, _controlCategoryId, _pointGroupCode, _controlPointId, _lackId, _points);
    setModified();

    if (_contributionCode == 5416) {
      updateBtsPoint(_pointGroupCode, _points);
    }
  }

  function addLacks(uint16[] _contributionCodes, int64[] _controlCategoryIds, uint16[] _pointGroupCodes, int64[] _controlPointIds, int64[] _lackIds, uint8[] _points) {
    //    require(msg.sender == inspectorAddress);
    for (uint16 i = 0; i < _contributionCodes.length; i++) {
      uint lacksIndex = numLacks++;
      newLacks[lacksIndex] = LackNew(_contributionCodes[i], _controlCategoryIds[i], _pointGroupCodes[i], _controlPointIds[i], _lackIds[i], _points[i]);
    }
    //  numLacks = _contributionCodes[0];
  }

  /* ==============================
    Calculate Direct Payment amount
    ===============================
    - All amounts in rappen */

  // 1110   Milchkühe
  // 1150   andere Kühe

  // 1128   weibliche Tiere über 365 - 730 Tage alt, ohne Abkalbung
  // 1141   weibliche Tiere über 160 - 365 Tage alt
  // 1142   weibliche Tiere bis 160 Tage alt (nur RAUS)
  // 1124   männliche Tiere, über 730 Tage alt
  // 1129   männliche Tiere, über 365 bis 730 Tage alt
  // 1143   männliche Tiere, über 160 bis 365 Tage alt
  // 1144   männliche Tiere, bis 160 Tage alt (nur RAUS)

  // For each pointGroup there is one CalcVariables
  uint16[9] pointGroupCodes = [1110, 1150, 1128, 1141, 1142, 1124, 1129, 1143, 1144];

  mapping (uint16 => CalcVariables) public pointGroups;

  struct CalcVariables {
  uint16 gve;
  uint16 points;
  uint16 btsPoints;
  uint256 btsTotal;
  uint256 btsDeduction;
  }

  function setGVE(uint16 _gve1110,
  uint16 _gve1150,
  uint16 _gve1128,
  uint16 _gve1141,
  uint16 _gve1142,
  uint16 _gve1124,
  uint16 _gve1129,
  uint16 _gve1143,
  uint16 _gve1144) {
    CalcVariables btsPointGroup = pointGroups[1110];
    btsPointGroup.gve = _gve1110;

    btsPointGroup = pointGroups[1150];
    btsPointGroup.gve = _gve1150;

    btsPointGroup = pointGroups[1128];
    btsPointGroup.gve = _gve1128;

    btsPointGroup = pointGroups[1141];
    btsPointGroup.gve = _gve1141;

    btsPointGroup = pointGroups[1142];
    btsPointGroup.gve = _gve1142;

    btsPointGroup = pointGroups[1124];
    btsPointGroup.gve = _gve1124;

    btsPointGroup = pointGroups[1129];
    btsPointGroup.gve = _gve1129;

    btsPointGroup = pointGroups[1143];
    btsPointGroup.gve = _gve1143;

    btsPointGroup = pointGroups[1144];
    btsPointGroup.gve = _gve1144;

    calculateBTS();
  }

  //  function getBTSGVE(uint16 _pointGroupCode) constant returns (uint16) {
  //    return pointGroups[_pointGroupCode].gve;
  //  }

  function updateBtsPoint(uint16 _pointGroupCode, uint16 _points) {
    CalcVariables btsPointGroup = pointGroups[_pointGroupCode];
    btsPointGroup.btsPoints = btsPointGroup.btsPoints + _points;
  }

  function calculateBTS() {
    for (uint16 i = 0; i < pointGroupCodes.length; i++) {
      CalcVariables btsPointGroup = pointGroups[pointGroupCodes[i]];
      //      btsPointGroup.btsTotal =  uint256(btsPointGroup.gve) * 9000;
      if ((pointGroupCodes[i] == 1142 || pointGroupCodes[i] == 1144) == false) {
        if (btsPointGroup.btsPoints < 110) {
          btsPointGroup.btsTotal = uint256(btsPointGroup.gve) * 9000;
          //btsPointGroup.btsDeduction = (btsPointGroup.btsPoints - 10) * (9000 / 100);
        }
      }
    }
  }

  function getFirstPaymentAmount() constant returns (uint256) {
    uint256 amount = 0;
    for (uint16 i = 0; i < pointGroupCodes.length; i++) {
      CalcVariables calcVar = pointGroups[pointGroupCodes[i]];
      amount = amount + calcVar.btsTotal;
    }
    return amount;
  }

}