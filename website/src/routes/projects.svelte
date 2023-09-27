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
	export let subtitle = '';
</script>

<card>
	<div
		class="transition-all group relative duration-300 ease-in-out from-s-blue rounded-xl shadow-2xl p-6 to-indigo-600 hover:shadow-[0_20px_50px_rgba(8,_112,_184,_0.7)] hover:shadow-s-blue/60"
	>
		<a {href} target="_blank">
			<div class="md:grid md:grid-flow-col">
				<div>
					<div class="col-span-4 w-full">
						<h1 class="text-sm lg:text-lg font-semibold">{duration}</h1>
						<div class="flex">
							<h1 class="font-bold text-white text-2xl lg:text-3xl pb-2">{title}</h1>
							<div
								class="w-4 lg:w-5 mt-2 ml-4 group-hover:-translate-y-2 group-hover:translate-x-2 duration-300 ease-in-out group-hover:text-white"
							>
								<FaLocationArrow />
							</div>
						</div>
						{#if subtitle != ''}
							<h2 class="text-sm lg:text-2xl pb-2 font-semibold">{subtitle}</h2>
						{/if}
						<p class="text-md md:text-lg lg:text-xl pr-4">{about}</p>
					</div>
				</div>
				<div class="col-span-1 my-auto pt-8">
					{#if img != ''}
						<div class="md:w-40 lg:w-64 rounded-xl">
							<img src={img} alt={title} class="rounded-lg shadow-xl" />
						</div>
					{/if}
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
