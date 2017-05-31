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

  int64 public userId;

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

  uint public numLacks;

  mapping (uint => Lack) public lacks;

  function Request(int64 _userId, uint16[] _contributionCodes, string _remark, address rbacAddress) public {
    rbac = RBAC(rbacAddress);
    //    m.sendToken(receiver, amount);
    userId = _userId;
    contributionCodes = _contributionCodes;
    remark = _remark;
    setCreated();
  }

  function setInspectorId(address _inspectorAddress){
    require(rbac.isAdmin(msg.sender));
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

  // For each pointGroup there is one CalcVariables
  uint16[10] pointGroupCodes = [1110, 1150, 1123, 1128, 1141, 1142, 1124, 1129, 1143, 1144];

  mapping (uint16 => CalcVariables) public btsPointGroups;

  struct CalcVariables {
  uint16 gve;
  uint16 points;
  uint16 total;
  uint16 deduction;
  }

  function updateBtsPoint(uint16 _pointGroupCode, uint16 _points) {
    CalcVariables btsPointGroup = btsPointGroups[_pointGroupCode];
    btsPointGroup.points = btsPointGroup.points + _points;
  }

  function calculateBTS() constant returns (uint16){
    // 1110    Milchkühe
    // 1150   andere Kühe
    // 1123   weibliche Tiere über 730 Tage alt, ohne Abkalbung
    // 1128    weibliche Tiere über 365 - 730 Tage alt, ohne Abkalbung
    // 1141    weibliche Tiere über 160 - 365 Tage alt
    // 1142   weibliche Tiere bis 160 Tage alt (nur RAUS)
    // 1124   männliche Tiere, über 730 Tage alt
    // 1129   männliche Tiere, über 365 bis 730 Tage alt
    // 1143   männliche Tiere, über 160 bis 365 Tage alt
    // 1144   männliche Tiere, bis 160 Tage alt (nur RAUS)

    uint16 sum = 0;

    for (uint16 i)

    // calculate deductions
    return sum;
  }

}


contract DirectPaymentCalculator {

  //function bts(uint _gve) {
  //
  //}

}



