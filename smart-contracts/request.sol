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

  uint public amountPreviousYear;

  RBAC rbac;

  struct Lack {
  uint16 contributionCode;
  int64 controlCategoryId;
  uint16 pointGroupCode;
  int64 controlPointId;
  int64 lackId;
  uint8 points;
  }

  uint public numLacks;

  mapping (uint => Lack) public lacks;


  function Request(uint16[] _contributionCodes, string _remark, address _rbacAddress, uint16[] _gves, uint _amountPreviousYear) public {
    rbac = RBAC(_rbacAddress);
    contributionCodes = _contributionCodes;
    remark = _remark;
    setGVE(_gves[0], _gves[1], _gves[2], _gves[3], _gves[4], _gves[5], _gves[6], _gves[7], _gves[8]);
    amountPreviousYear = _amountPreviousYear;
    setCreated();
  }

  function setInspectorId(address _inspectorAddress){
    require(rbac.isAdmin(msg.sender) || rbac.isCantonEmployee(msg.sender));
    require(rbac.isInspector(_inspectorAddress));
    inspectorAddress = _inspectorAddress;
    setModified();
  }

  function addLacks(uint16[] _contributionCodes, int64[] _controlCategoryIds, uint16[] _pointGroupCodes, int64[] _controlPointIds, int64[] _lackIds, uint8[] _points) {
    require(msg.sender == inspectorAddress);
    for (uint16 i = 0; i < _contributionCodes.length; i++) {
      uint lacksIndex = numLacks++;
      lacks[lacksIndex] = Lack(_contributionCodes[i], _controlCategoryIds[i], _pointGroupCodes[i], _controlPointIds[i], _lackIds[i], _points[i]);
      if (_contributionCodes[i] == 5416) {
        updateBtsPoint(_pointGroupCodes[i], _points[i]);
      }
      if (_contributionCodes[i] == 5417) {
        updateRausPoint(_pointGroupCodes[i], _points[i]);
      }
    }
    setModified();
  }


  // For each pointGroup there is a PointGroupCalculation
  uint16[9] pointGroupCodes = [1110, 1150, 1128, 1141, 1142, 1124, 1129, 1143, 1144];


  struct PointGroupCalculation {
  uint16 gve;
  uint16 btsPoints;
  uint256 btsTotal;
  uint256 btsDeduction;
  uint16 rausPoints;
  uint256 rausTotal;
  uint256 rausDeduction;
  }

  mapping (uint16 => PointGroupCalculation) public pointGroups;

  function setGVE(uint16 _gve1110, uint16 _gve1150, uint16 _gve1128, uint16 _gve1141, uint16 _gve1142, uint16 _gve1124, uint16 _gve1129, uint16 _gve1143, uint16 _gve1144) {
    PointGroupCalculation storage btsPointGroup = pointGroups[1110];
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
    setModified();
  }


  function updateBtsPoint(uint16 _pointGroupCode, uint16 _points) {
    PointGroupCalculation storage pointGroupCalculation = pointGroups[_pointGroupCode];
    pointGroupCalculation.btsPoints = pointGroupCalculation.btsPoints + _points;
  }

  function updateRausPoint(uint16 _pointGroupCode, uint16 _points) {
    PointGroupCalculation storage pointGroupCalculation = pointGroups[_pointGroupCode];
    pointGroupCalculation.rausPoints = pointGroupCalculation.rausPoints + _points;
  }

  function calculateBTS() {
    for (uint16 i = 0; i < pointGroupCodes.length; i++) {
      PointGroupCalculation storage btsPointGroup = pointGroups[pointGroupCodes[i]];

      if ((pointGroupCodes[i] == 1142 || pointGroupCodes[i] == 1144) == false) {
        btsPointGroup.btsTotal = uint256(btsPointGroup.gve) * 9000;
        if (btsPointGroup.btsPoints == 0) {
          continue;
        }
        if (btsPointGroup.btsPoints > 110) {
          btsPointGroup.btsDeduction = btsPointGroup.btsTotal;
          continue;
        }
        btsPointGroup.btsDeduction = (btsPointGroup.btsPoints - 10) * (9000 / 100);
      }
    }
  }

  function calculateRAUS() {
    for (uint16 i = 0; i < pointGroupCodes.length; i++) {
      PointGroupCalculation storage rausPointGroup = pointGroups[pointGroupCodes[i]];
      uint256 multiplier = 0;
      if (pointGroupCodes[i] == 1142 || pointGroupCodes[i] == 1144) {
        multiplier = 37000;
      }
      else {
        multiplier = 19000;
      }
      rausPointGroup.rausTotal = uint256(rausPointGroup.gve) * multiplier;
      if (rausPointGroup.rausPoints == 0) {
        continue;
      }
      if (rausPointGroup.rausPoints > 110) {
        rausPointGroup.rausDeduction = rausPointGroup.rausTotal;
        continue;
      }
      rausPointGroup.rausDeduction = (rausPointGroup.rausPoints - 10) * (multiplier / 100);

    }
  }

  function getFirstPaymentAmount() constant returns (uint256) {
    uint256 amount = 0;
    if (amountPreviousYear > 0) {
      // first payment is 50% of amount of previous year
      amount = amountPreviousYear / 2;
    }
    return amount;
  }

  function getFinalPaymentAmount() constant returns (uint256){

    uint256 amount = 0;
    for (uint16 i = 0; i < pointGroupCodes.length; i++) {
      PointGroupCalculation storage pointGroup = pointGroups[pointGroupCodes[i]];
      amount += (pointGroup.btsTotal - pointGroup.btsDeduction);
      amount += (pointGroup.rausTotal - pointGroup.rausDeduction);
    }
    return amount - getFirstPaymentAmount();
  }
}