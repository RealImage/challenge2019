import sqlite3  from 'sqlite3'
import { open } from 'sqlite'
const sql=sqlite3.verbose();
const db=new sqlite3.Database(':memory:')

export default db;
