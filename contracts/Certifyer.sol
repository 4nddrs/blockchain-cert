// // SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract Certifyer {
  address public admin;

  //Mapping to save: Documents and their hash => true is valid
  mapping(bytes32 => bool) public certificates;

  //Event to say to the world that a document has been certified
  event CertificateCreated(bytes32 indexed datahash, uint256 timestamp);

  constructor(){
    admin = msg.sender;
  }

  //Just the admin can certify documents
  function registerCertificate(bytes32 datahash) public {
    require(msg.sender == admin, "Only the admin can certify documents");
    certificates[datahash] = true;
    emit CertificateCreated(datahash, block.timestamp);
  }

  //Anyone can check if a document is certified by its hash
  function validateCertificate(bytes32 datahash) public view returns (bool) {
    return certificates[datahash];
  }
}
