<script>
	// @ts-nocheck

	// https://github.com/topojson/us-atlas
	// https://github.com/d3/d3-geo
	// https://observablehq.com/@mbostock/u-s-state-map
	// TODO: https://observablehq.com/@d3/u-s-map
	// TODO: https://observablehq.com/@jeantimex/us-state-county-map

	import { onMount } from 'svelte';
	import * as topojson from 'topojson-client';
	import { geoPath, geoEqualEarth } from 'd3-geo';
	import { draw } from 'svelte/transition';
	import { countries, temperatures } from './example.json';

	let loading_map = false;

	const projection = geoEqualEarth();
	const path = geoPath(projection);

	// let states = [];
	// let counties = [];
	let country_features = [];
	// let country_arr = [];
	let mesh;
	let selected;
	//$: console.log({ selected })

	onMount(async () => {
		const globe = await fetch('https://cdn.jsdelivr.net/npm/world-atlas@2/countries-50m.json').then(
			(d) => d.json()
		);
		country_features = topojson.feature(globe, globe.objects.countries).features;
		loading_map = true;
	});
</script>

{#if loading_map}
	<div class="border-white shadow-2xl">
		<svg viewBox="400 40 250 150">
			{#each country_features as feature, i}
				<!-- {#if countries[feature.id] > 0} -->
				<g fill={countries[feature.id] ?? 'white'}>
					<path
						d={path(feature)}
						on:click={() => (selected = feature)}
						class="state"
						in:draw={{ delay: i * 50, duration: 1000 }}
					/>
				</g>
			{/each}
		</svg>
	</div>
{:else}
	<button type="button" class="flex mx-auto pt-32" disabled>
		<svg
			stroke="currentColor"
			fill="none"
			stroke-width="0"
			viewBox="0 0 24 24"
			height="1em"
			width="1em"
			class="animate-spin my-auto"
			xmlns="http://www.w3.org/2000/svg"
			><path
				d="M12 22C17.5228 22 22 17.5228 22 12H19C19 15.866 15.866 19 12 19V22Z"
				fill="currentColor"
			/><path
				d="M2 12C2 6.47715 6.47715 2 12 2V5C8.13401 5 5 8.13401 5 12H2Z"
				fill="currentColor"
			/></svg
		>
		<h1 class="pl-2">Processing Map Data</h1>
	</button>{/if}

<div class="selectedName">{selected?.properties.name ?? ''}</div>

<style>
	.state:hover {
		fill: hsl(0 0% 50% / 20%);
	}

	.selectedName {
		text-align: center;
		margin-top: 8px;
		font-size: 1.5rem;
	}
</style>
