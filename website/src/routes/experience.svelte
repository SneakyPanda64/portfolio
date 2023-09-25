<script>
	// @ts-nocheck

	// @ts-ignore
	import FaLocationArrow from 'svelte-icons/fa/FaLocationArrow.svelte';
	import Skills from './skills.svelte';
	export let role = 'Web Developer';
	export let company = {
		title: 'Talon Outdoors',
		website: 'https://google.com'
	};
	export let duration = '2018 — PRESENT';
	export let level = 'Internship';
	export let about =
		'Deliver high-quality, robust production code for a diverse array of projects for clients including Harvard Business School, Everytown for Gun Safety, Pratt Institute, Koala Health, Vanderbilt University, The 19th News, and more. Provide leadership within engineering department through close collaboration, knowledge shares, and mentorship.';
	export let skills = ['Next.js', 'Jest', 'Sass', 'Javascript', 'Typescript'];
</script>

<div class="pt-16">
	<card>
		<a href={company.website} target="_blank" class="">
			<div
				class="transition-all group relative duration-300 shadow-2xl ease-in-out from-s-blue rounded-xl p-6 to-indigo-600 hover:shadow-[0_20px_50px_rgba(8,_112,_184,_0.7)] hover:shadow-s-blue/60"
			>
				<div class="">
					<!-- <h1 class="font-semibold text-2xl">2023</h1> -->
					<h1 class="text-sm font-semibold">{duration}</h1>
				</div>
				<div class="">
					<div class="flex">
						<h1 class="text-2xl font-bold text-white">{role} • {company.title}</h1>
						<div
							class="w-4 mt-2 ml-4 group-hover:-translate-y-2 group-hover:translate-x-2 duration-300 ease-in-out group-hover:text-white"
						>
							<FaLocationArrow />
						</div>
					</div>
					<h2 class="text-xl font-semibold">{level}</h2>
					<p class="">
						{about}
					</p>
					<Skills {skills} />
				</div>
			</div></a
		>
	</card>
</div>
