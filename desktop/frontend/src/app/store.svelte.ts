import { IsStoreLocked, UnlockStore, ChangePassword, ReadTest, WriteTest, IsFirstRun, FirstRunCompleted, IsStoreProtected } from '@go';

let locked : boolean = $state(true);
let firstRun : boolean = $state(true);
let _protected : boolean = $state(false);

export async function initialize() {
  locked = await IsStoreLocked();
  firstRun = await IsFirstRun();
  _protected = await IsStoreProtected();
}

export async function unlock(password: string = "", onunlock : () => void = () => {}) {
  if (!locked) {
    return;
  }
  await UnlockStore(password);
  locked = await IsStoreLocked();
  if (!locked) {
    onunlock();
  }
  _protected = await IsStoreProtected();
}

export async function changePassword(password: string = "") {
  if (locked) {
    return;
  }
  await ChangePassword(password);
   _protected = await IsStoreProtected();
}

async function readTest() {
  return await ReadTest();
}

async function writeTest(value: string) {
  return await WriteTest(value);
}

export default {
  get locked() {
    return locked;
  },
  get protected() {
    return _protected;
  },
  readTest,
  writeTest,
  changePassword,
  unlock,
  get firstRun() {
    return firstRun;
  },
  firstRunCompleted: async () : Promise<void> => {
    await FirstRunCompleted();
    firstRun = await IsFirstRun();
  }
}