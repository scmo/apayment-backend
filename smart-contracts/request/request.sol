pragma solidity ^0.4.0;

contract mortal {
  /* Define variable owner of the type address*/
  address owner;

  /* this function is executed at initialization and sets the owner of the contract */
  function mortal() { owner = msg.sender; }

  /* Function to recover the funds on the contract */
  function kill() { if (msg.sender == owner) selfdestruct(owner); }
}

contract Request is mortal {


  int64 public userId;

  int64 public inspectorId;

  uint16[] public contributionCodes;

  string public remark;

  function Request(int64 _userId, uint16[] _contributionCodes, string _remark) public {
    userId = _userId;
    contributionCodes = _contributionCodes;
    remark = _remark;
  }

  function createdTimestamp() constant returns (uint){
    return block.timestamp;
  }

  function setInspectorId(int64 _inspectorId){
    inspectorId = _inspectorId;
  }

}