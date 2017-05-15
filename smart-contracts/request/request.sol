pragma solidity ^0.4.11;


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

  struct Lack {
  uint16 contributionCode;
  string controlCategoryId;
  string pointGroupId;
  string controlPointId;
  int64 lackId;
  }

  uint public numLacks;

  mapping (uint => Lack) public lacks;

  function Request(int64 _userId, uint16[] _contributionCodes, string _remark) public {
    userId = _userId;
    contributionCodes = _contributionCodes;
    remark = _remark;

    setCreated();
  }

  function setInspectorId(address _inspectorAddress){
    inspectorAddress = _inspectorAddress;
    setModified();
  }


  function addLack(uint16 _contributionCode, string _controlCategoryId, string _pointGroupId, string controlPointId, int64 lackId) {
    require(tx.origin == inspectorAddress);
    uint lacksIndex = numLacks++;
    lacks[lacksIndex] = Lack(_contributionCode, _controlCategoryId, _pointGroupId, controlPointId, lackId);
    setModified();
  }

}