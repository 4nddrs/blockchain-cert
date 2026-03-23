// // SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract Certifyer {
  
  address public owner;

  struct Certificate{
    string studentName;
    string courseName;
    string issuerName;
    uint256 dateEmited;
    bool isValid;
  }

  //Mapping to save: Documents and their hash => true is valid
  mapping(bytes32 => Certificate) public certificates;

  // Control who can emit certificates
  mapping(address => bool) public authorizedIssuers;

  //Event to say to the world that a document has been certified
  event CertificateCreated(bytes32 indexed datahash, string studentName, string issuer);
  event IssuerAuthorized(address indexed issuer);

  constructor(){
    owner = msg.sender;
    authorizedIssuers[msg.sender] = true; // The deployer is an authorized issuer
  }

  modifier onlyOwner() {
    require(msg.sender == owner, "Only the owner can perform this action");
    _;
  }

  // BUSSINESS FUNCTION: Add an authorized issuer
  function addAuthorizedIssuer(address issuer) public onlyOwner {
    authorizedIssuers[issuer] = true;
    emit IssuerAuthorized(issuer);
  }
  
  // BUSSINESS FUNCTION: Save all the information of the certificate
  function registerCertificate(

    bytes32 datahash,
    string memory studentName,
    string memory courseName,
    string memory issuerName
                              
  ) public {
      require(authorizedIssuers[msg.sender], "Only authorized issuers can register certificates");
      require(!certificates[datahash].isValid, "Certificate with this hash already exists");
      
      certificates[datahash] = Certificate({
        studentName: studentName,
        courseName: courseName,
        issuerName: issuerName,
        dateEmited: block.timestamp,
        isValid: true
      });

      emit CertificateCreated(datahash, studentName, issuerName);

  }

  //Anyone can check if a document is certified by its hash
  function validateCertificate(bytes32 datahash) public view returns (Certificate memory) {
    return certificates[datahash];
  }
}
