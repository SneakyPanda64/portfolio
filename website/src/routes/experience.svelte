<script>
	import FaLocationArrow from 'svelte-icons/fa/FaLocationArrow.svelte';
	let skills = ['1qwdwq', 'ewfew', '32423', 'wiiwfwf', 'wdewifewf'];
	import { flip } from 'svelte/animate';
	import { quintOut } from 'svelte/easing';
</script>

<card>
	<a href="https://google.com"
		><div
			class="transition-all group relative duration-300 ease-in-out from-s-blue rounded-xl p-6 my-12 to-indigo-600 hover:shadow-[0_20px_50px_rgba(8,_112,_184,_0.7)] mt-12 hover:shadow-cyan-500/40"
		>
			<div
				class="absolute right-0 p-0 pr-6 w-12 group-hover:-translate-y-2 group-hover:translate-x-2 duration-300 ease-in-out group-hover:text-white"
			>
				<FaLocationArrow />
			</div>
			<h1 class="text-2xl font-bold text-white">Lead Engineer - Upstatement</h1>
			<h2 class="text-xl font-semibold">Senior Engineer</h2>
			<p class="">
				Deliver high-quality, robust production code for a diverse array of projects for clients
				including Harvard Business School, Everytown for Gun Safety, Pratt Institute, Koala Health,
				Vanderbilt University, The 19th News, and more. Provide leadership within engineering
				department through close collaboration, knowledge shares, and mentorship.
			</p>
			<ul class="flex gap-x-3 pt-6">
				{#each skills as skill}
					<li class="bg-s-blue text-s-blue bg-opacity-10 font-bold p-2 rounded-full px-4 text-sm">
						{skill}
					</li>
				{/each}
			</ul>
		</div></a
	>
</card>
