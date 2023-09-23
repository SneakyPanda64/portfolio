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
	import { wsConnection } from './stores/ws';
	import { heatmap } from './stores/heatmap';
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
		$wsConnection.send(
			JSON.stringify({
				action: 'stats',
				value: ''
			})
		);
	});
</script>

{#if loading_map && Object.keys($heatmap).length > 0}
	<div>
		<div class="border-white shadow-2xl">
			<svg viewBox="400 20 250 150">
				{#each country_features as feature, i}
					<!-- {@const hash =(feature.id))} -->
					<g fill={$heatmap[parseInt(feature.id)] ?? 'white'}>
						<path
							d={path(feature)}
							on:click={() => (selected = feature)}
							class="country"
							in:draw={{ delay: i * 50, duration: 1000 }}
						/>
					</g>
				{/each}
			</svg>
			<div class="flex py-4 pl-2">
				<!-- <h1 class="pr-2 font-semibold text-white">Least</h1> -->
				{#each ['#ffffff', '#fffdc4', '#fffb80', '#ffe14a', '#ffcc4a', '#ffbd4a', '#ffa44a', '#ff8c4a', '#ff6e4a', '#ff564a'] as colour}
					<div class={`h-4 my-auto w-4 mr-1`} style={`background-color: ${colour};`} />
				{/each}
				<!-- <h1 class="pr-2 font-semibold text-white">Most</h1> -->
			</div>
		</div>
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
	.country:hover {
		fill-opacity: 95%;
	}

	.selectedName {
		text-align: center;
		margin-top: 8px;
		font-size: 1.5rem;
	}
</style>
