<script lang="ts">
  import { _ } from "svelte-i18n";
  import ColorPicker from 'svelte-awesome-color-picker';
  import type { RgbaColor } from 'svelte-awesome-color-picker';
  import { color as col } from '@models';

  interface Props {
    label: string;
    color: col.RGBA;
  }

  let { label, color = $bindable() } : Props = $props();

  let rgb : RgbaColor = $state({
    r: color.R,
    g: color.G,
    b: color.B,
    a: color.A / 255,
  });

  $effect(() => {
    if (rgb) {
      color.R = Math.round(rgb.r * rgb.a);
      color.G = Math.round(rgb.g * rgb.a);
      color.B = Math.round(rgb.b * rgb.a);
      color.A = Math.round(rgb.a * 255);
    }
  });

</script>

<ColorPicker 
  {label}
  bind:rgb
  texts={{
		label: {
			h: $_("colorpicker.label.h"),
			s: $_("colorpicker.label.s"),
			v: $_("colorpicker.label.v"),
			r: $_("colorpicker.label.r"),
			g: $_("colorpicker.label.g"),
			b: $_("colorpicker.label.b"),
			a: $_("colorpicker.label.a"),
			hex: $_("colorpicker.label.hex"),
			withoutColor: $_("colorpicker.label.withoutColor"),
		},
		color: {
			rgb: $_("colorpicker.color.rgb"),
			hsv: $_("colorpicker.color.hsv"),
			hex: $_("colorpicker.color.hex"),
		},
		changeTo: $_("colorpicker.changeTo"),
	}}
/>