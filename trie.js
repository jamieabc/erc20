const Trie = require('merkle-patricia-tree/secure');
const levelup = require('levelup');
const leveldown = require('leveldown');
const assert = require('assert');
const utils = require('ethereumjs-util');
const BN = utils.BN;
const Account = require('ethereumjs-account');

const dbPath = '/Users/Aaron/Documents/Workspace/solidity/erc20/eth/geth/chaindata';
const db = levelup(leveldown(dbPath));

const root = '0x923cecf6550b6c478e9538a890571545de7d6570cb6a159b11e759eaf7e4d81b';
const contractAddr= '0x06c5d911C9F5e526473331086B31d0b58FBd916d';

const trie = new Trie(db, root);

trie.get(contractAddr, (err, raw) => {
  const account = new Account(raw);
  const storageRoot = account.stateRoot;
  const storageTrie = new Trie(db, storageRoot);
  const stream = storageTrie.createReadStream();

  stream.on('data', (data) => null);
});
