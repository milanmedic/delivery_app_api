#!/usr/bin/env node
'use strict';

const fs = require('fs');
const path = require('path');
const pkg = require(path.join(__dirname,'config.json'));

const commitRegExp = new RegExp(pkg.config.commitMessageRegex);

const errorMsg = pkg.config.commitMessageComment;

try {
    const commitMsg = fs.readFileSync(process.argv[2], 'utf8');
    if (!commitRegExp.test(commitMsg)) {
        console.log(errorMsg);
        process.exit(1);
    }
} catch (err) {
    console.log('[GUARD]: Error: Commit message file not found in git repo ', err);
}
