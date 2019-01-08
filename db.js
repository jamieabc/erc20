const levelup = require('levelup');
const leveldown = require('leveldown');
const dump = require('level-dump');

const dbPath = '/Users/Aaron/Documents/Workspace/solidity/erc20/eth/geth/chaindata';
const db = levelup(leveldown(dbPath));

dump(db);
// dump.allKeys(db);
// dump.allValues(db);
// dump.allEntries(db);
