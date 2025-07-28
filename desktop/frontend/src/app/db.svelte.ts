import { IsLocked, OpenDB, ChangePassword, ReadTest, WriteTest } from '@go';
import {initialize as initSettings, default as settings} from './settings.svelte';

let locked : boolean = $state(true);
let encrypted : boolean = $state(false);

export async function initialize() {
  locked = await IsLocked();
  encrypted = settings.settings.enc !== "";
  console.log("db", $state.snapshot(locked), $state.snapshot(encrypted));
  if (locked && !encrypted) {
    await unlock();
  }
}

export async function unlock(password: string = "") {
  if (!locked) {
    return;
  }
  await OpenDB(password);
  await initSettings();
  locked = await IsLocked();
  encrypted = settings.settings.enc !== "";
}

export async function changePassword(password: string = "") {
  if (locked) {
    return;
  }
  await ChangePassword(password)
  await initSettings();
  encrypted = settings.settings.enc !== "";
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
  get encrypted() {
    return encrypted;
  },
  readTest,
  writeTest,
  changePassword,
  unlock,
}