const Trie = require('merkle-patricia-tree/secure');
const levelup = require('levelup');
const leveldown = require('leveldown');
const assert = require('assert');
const utils = require('ethereumjs-util');
const BN = utils.BN;
const Account = require('ethereumjs-account');

const dbPath = '/Users/Aaron/Documents/Workspace/solidity/erc20/eth/geth/chaindata';
const db = levelup(leveldown(dbPath));

const stateRoots = {
  contract: '0x1eead80a3a355c48c005c656063355ed36bda3a166b95cf693280d5f147cbc2c',
  txNewAcc1: '0x63ce1835163ef2f688ca234297b03da70fa788ff20888dba77829bcc48d10301',
  txNewAcc2: '0xc7de53dbfee29d8e8b09542c41f527d6e7e5b7a1ed3a957f450aac62e68395cd',
  txOldAcc1: '0x1d10fae89220395c080dea426186b3c17b647703c8719fb393d07ba5fdbd08d9',
  txOldAcc2: '0xaad7da537bccea49ffbddb85b0df31cfed237cabcadeaf086d946a7892e9ec5c'
};

const contractAddr= '0xA111E8cA6Df293EA535b23f9335Ff55A2066B7c9';

Object.keys(stateRoots).forEach((key) => {
  const stateRoot = stateRoots[key];
  const trie = new Trie(db, stateRoot);

  trie.get(contractAddr, (err, raw) => {
    const account = new Account(raw);
    const storageRoot = account.stateRoot;
    const storageTrie = new Trie(db, storageRoot);
    const stream = storageTrie.createReadStream();

    stream.on('data', ()=> null);
    stream.on('end', () => console.log(key));
  });
});
