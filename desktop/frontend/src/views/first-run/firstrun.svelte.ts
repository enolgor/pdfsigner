
let step : number = $state(0);
let steps : number = $state(0);
export default {
  set steps(v : number) {
    steps = v
  },
  get step() {
    return step;
  },
  get done() {
    return step >= steps;
  },
  advance: () => {
    step += 1;
  },
}