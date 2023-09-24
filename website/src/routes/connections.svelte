<script>
	export let icon;
	export let href;
</script>

<a
	{href}
	target="_blank"
	class="w-12 h-12 p-2.5 rounded-xl bg-s-blue bg-opacity-10 text-s-blue hover:bg-opacity-20"
>
	<svelte:component this={icon} />
	<!-- <this:svelsvelte:component></this:svelsvelte:component> -->
</a>
