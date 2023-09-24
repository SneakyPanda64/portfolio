<script>
	export let linkName = '';
</script>

<a href={`#${linkName}`} class="group cursor-pointer">
	<div class="w-1/2 group-hover:w-3/4 flex transition-all duration-300">
		<div
			class="h-[0.1rem] my-auto bg-s-blue w-full bg-opacity-50 group-hover:bg-opacity-100 mr-4"
		/>
		<h1 class="text-lg font-semibold group-hover:text-white">{linkName.toUpperCase()}</h1>
	</div>
</a>
