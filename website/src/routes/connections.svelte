<script>
	export let icon;
	export let href;
</script>

<div class="w-12 h-12 p-2.5 rounded-xl bg-s-blue bg-opacity-10 text-s-blue hover:bg-opacity-20">
	<a {href} target="_blank"><svelte:component this={icon} /></a>
	<!-- <this:svelsvelte:component></this:svelsvelte:component> -->
</div>
