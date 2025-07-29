import { IsStoreLocked, UnlockStore, ChangePassword, ReadTest, WriteTest } from '@go';

let locked : boolean = $state(true);

export async function initialize() {
  locked = await IsStoreLocked();
  console.log("store is locked", $state.snapshot(locked))
}

export async function unlock(password: string = "") {
  if (!locked) {
    return;
  }
  await UnlockStore(password);
  locked = await IsStoreLocked();
}

export async function changePassword(password: string = "") {
  if (locked) {
    return;
  }
  await ChangePassword(password);
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
  readTest,
  writeTest,
  changePassword,
  unlock,
}