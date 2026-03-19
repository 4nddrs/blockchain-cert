// // SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract Certificador {
  address public admin;

  //Mapping to save: Documents and their hash => true is valid
  mapping(bytes32 => bool) public certificados;

  //Event to say to the world that a document has been certified
  event CertificadoCreado(bytes32 indexed datahash, uint256 timestamp);

  constructor(){
    admin = msg.sender;
  }

  //Just the admin can certify documents
  function registrarCertificado(bytes32 datahash) public {
    require(msg.sender == admin, "Only the admin can certify documents");
    certificados[datahash] = true;
    emit CertificadoCreado(datahash, block.timestamp);
  }

  //Anyone can check if a document is certified by its hash
  function validarCertificado(bytes32 datahash) public view returns (bool) {
    return certificados[datahash];
  }
}
