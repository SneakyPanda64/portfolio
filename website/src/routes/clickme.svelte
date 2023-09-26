<script>
	// @ts-nocheck

	import { showStats } from './stores/stats';
	import { wsConnection } from './stores/ws';
	import { clicks } from './stores/clicks';
	// let clicks = 0;
	let clicked = false;

	const clickEvent = () => {
		// if (clicked) return;
		// clicks += 1;
		clicked = true;
		if ($wsConnection !== undefined) {
			// @ts-ignore
			$wsConnection.send(
				JSON.stringify({
					action: 'click',
					value: ''
				})
			);
		}
	};
	const showStatsHandler = () => {
		console.log('changing show stats var');
		$showStats = !$showStats;
	};
</script>

<div class="">
	<div class="flex">
		<button
			on:click={clickEvent}
			class={'bg-s-blue flex-grow-0 py-3 px-6 lg:py-5 lg:px-10 rounded-xl bg-opacity-30 text-s-blue text-xl lg:text-3xl font-semibold transition-all duration-300 hover:bg-opacity-40 hover:text-white'}
		>
			{#if clicked}
				I have been clicked {$clicks} times!
			{:else}
				Click Me!
			{/if}
		</button>
		{#if clicked}
			<button
				on:click={showStatsHandler}
				class="bg-s-blue bg-opacity-10 text-s-blue rounded-xl ml-2 px-4 font-semibold hover:bg-opacity-30 hover:text-white"
				>{$showStats ? 'Hide' : 'Show'} statistics</button
			>
		{/if}
	</div>
</div>
