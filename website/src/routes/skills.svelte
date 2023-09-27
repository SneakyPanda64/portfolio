<script>
	export let skills = [];
</script>

<ul class="flex flex-wrap gap-x-3 pt-6 gap-y-2">
	{#each skills as skill}
		<li
			class="bg-s-blue text-s-blue bg-opacity-10 font-bold p-2 rounded-full px-4 text-sm md:text-lg lg:text-xl flex"
		>
			<h1 class="my-auto">{skill}</h1>
		</li>
	{/each}
</ul>
