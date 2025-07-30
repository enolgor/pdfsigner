export enum views {
  settings = 'settings',
  certificates = 'certificates',
  stamps = 'stamps',
  sign = 'sign',
  loading = 'loading',
  help = 'help',
}

let _view : views = $state(views.loading);

export const controller = {
  get view() : views {
    return _view;
  },
  set view(value: views) {
    _view = value;
  }
};