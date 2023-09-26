<script>
	export let icon;
	export let href;
</script>

<div
	class="w-20 h-20 p-4 md:p-2.5 md:w-12 lg:p-4 md:h-12 lg:w-20 lg:h-20 rounded-xl bg-s-blue bg-opacity-10 text-s-blue hover:bg-opacity-20"
>
	<a {href} target="_blank"><svelte:component this={icon} /></a>
	<!-- <this:svelsvelte:component></this:svelsvelte:component> -->
</div>
