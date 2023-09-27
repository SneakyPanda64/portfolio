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
		if ($wsConnection !== undefined && $clicks['disabled'] == false) {
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
			class={'flex-grow-0 py-3 px-6 lg:py-5 lg:px-10 rounded-xl bg-opacity-30  text-xl lg:text-3xl font-semibold transition-all duration-300  ' +
				($clicks['disabled']
					? 'bg-red-500 text-red-500'
					: 'bg-s-blue text-s-blue hover:bg-opacity-40 hover:text-white')}
		>
			{#if $wsConnection === undefined || $wsConnection.readyState !== WebSocket.OPEN}<div
					class="flex"
				>
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
					<h1 class="pl-2">Connecting</h1>
				</div>{:else if clicked}
				{#if $clicks['disabled']}
					Rate limited
				{:else}
					I have been clicked {$clicks['clicks']} times!
				{/if}
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
