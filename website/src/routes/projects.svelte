<script>
	import FaLocationArrow from 'svelte-icons/fa/FaLocationArrow.svelte';
	import Connections from './connections.svelte';
	import Skills from './skills.svelte';
	export let title = '';
	export let about = '';
	export let img = '';
	export let duration = '';
	export let connections = [];
	export let href = '';
	export let skills = [];
</script>

<card>
	<div
		class="transition-all group relative duration-300 ease-in-out from-s-blue rounded-xl shadow-2xl p-6 to-indigo-600 hover:shadow-[0_20px_50px_rgba(8,_112,_184,_0.7)] hover:shadow-s-blue/60"
	>
		<a {href} target="_blank">
			<div class="md:grid md:grid-flow-col">
				<div>
					<div class="col-span-4 w-full">
						<h1 class="text-sm font-semibold">{duration}</h1>
						<div class="flex">
							<h1 class="font-bold text-white text-2xl pb-2">{title}</h1>
							<div
								class="w-4 mt-2 ml-4 group-hover:-translate-y-2 group-hover:translate-x-2 duration-300 ease-in-out group-hover:text-white"
							>
								<FaLocationArrow />
							</div>
						</div>
						<p class="text-md pr-4">{about}</p>
					</div>
				</div>
				<div class="col-span-1 my-auto pt-8">
					<div class="md:w-40 rounded-xl">
						<img src={img} alt={title} class="rounded-lg shadow-xl" />
					</div>
					<!-- <div class="w-16 h-16 bg-red-500"> -->
					<!-- <FaGithub /> -->
					<ul class="pt-6 flex justify-between">
						{#each connections as connection}
							<Connections href={connection.href} icon={connection.icon} />
						{/each}
					</ul>

					<!-- </div> -->
				</div>
			</div>
		</a>
		<Skills {skills} />
	</div>
</card>
