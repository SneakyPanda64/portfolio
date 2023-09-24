<script>
	// @ts-nocheck
	import { showStats } from './stores/stats';
	import { wsConnection } from './stores/ws';
	import { clicks } from '../routes/stores/clicks';
	import FaGithub from 'svelte-icons/fa/FaGithub.svelte';
	import FaLinkedinIn from 'svelte-icons/fa/FaLinkedinIn.svelte';
	import FaGooglePlay from 'svelte-icons/fa/FaGooglePlay.svelte';
	import LinkComponent from './link.svelte';
	import Experience from './experience.svelte';
	import Clickme from './clickme.svelte';
	import Map from './map.svelte';
	import { onMount } from 'svelte';
	import { heatmap } from './stores/heatmap';
	import Connections from './connections.svelte';
	// import { env } from '$lib/env';
	// const secret = env.YOUR_SECRET;
	// @ts-ignore

	function followMouse(event) {
		console.log('Follow');
		let elem = document.querySelector('#cursor');
		if (elem !== null) {
			let y = event.pageY - elem.offsetHeight / 2;
			let x = event.pageX - elem.offsetWidth / 2;
			elem.style.setProperty('position', 'absolute');
			elem.style.setProperty('left', x + 'px');
			elem.style.setProperty('top', y + 'px');
		}
	}
	onMount(async () => {
		function connect() {
			const ws = new WebSocket('ws://localhost:3000/ws/123?v=1.0');
			ws.addEventListener('open', () => {
				console.log('connected');
				$wsConnection = ws;
			});
			ws.addEventListener('message', (message) => {
				// Parse the incoming message here
				if (message.data !== '') {
					const data = JSON.parse(message.data);
					console.log('recieved', data);
					if (data.type === 'clicks') {
						$clicks = data.value;
					} else if (data.type === 'stats') {
						$heatmap = JSON.parse(data.value);
						console.log($heatmap);
					}
				} else {
					console.log('recieved', message);
				}
			});
			ws.addEventListener('close', () => {
				console.log('closed connection');
				setTimeout(function () {
					connect();
				}, 1000);
			});
		}
		connect();
	});
</script>

<!-- svelte-ignore a11y-no-static-element-interactions -->
<div id="main" class="relative" on:mousemove={followMouse}>
	<div
		id="cursor"
		class="bg-gradient-to-r from-s-purple to-s-blue w-32 h-32 absolute bg-blend-multiply blur-3xl opacity-60 -z-50"
	/>
	<div class="min-h-screen w-screen pt-24 px-32 text-s-light-gray" on:wheel={followMouse}>
		<div class="grid grid-cols-2 gap-x-12">
			<div class="fixed w-1/2">
				<div>
					<h1 class="text-5xl text-white font-bold">ALEXANDER HEATHER</h1>
					<h2 class="text-3xl text-white font-medium text-opacity-90 pb-6 pt-2">WEB DEVELOPER</h2>
					<h3 class="text-2xl font-medium w-2/3">
						I excel in crafting high-performance websites that seamlessly integrate stunning visual
						elements
					</h3>
				</div>
				<div class="pt-16 w-1/2">
					<LinkComponent linkName="about" />
					<LinkComponent linkName="experiences" />
					<LinkComponent linkName="projects" />
				</div>
				<div class="pt-16">
					<Clickme />
				</div>
				<div class="flex gap-x-4 pt-12">
					<Connections icon={FaGithub} href={'https://github.com/SneakyPanda64'} />
					<Connections
						icon={FaLinkedinIn}
						href={'https://www.linkedin.com/in/alexander-heather-b98832249/'}
					/>
					<Connections
						icon={FaGooglePlay}
						href={'https://play.google.com/store/apps/developer?id=AtlasServers'}
					/>
				</div>
			</div>
			<div />
			<div>
				<!-- <Map /> -->
				{#if $showStats}
					<Map />
				{:else}
					<div>
						<p class="text-lg font-normal">
							Commodo voluptate irure magna laboris sint nulla. Do enim duis ea dolor est. Lorem
							dolore excepteur dolor <b>qui</b> ut officia aliqua commodo ipsum. Amet commodo
							nostrud in excepteur <b>qui</b>
							sit <b>non</b> <b>qui</b>s tempor nostrud tempor ipsum. Cillum magna ad esse
							reprehenderit. Culpa veniam mollit <b>non</b> cillum cillum est <b>qui</b> velit.
							Cupidatat sint exercitation excepteur duis voluptate magna veniam labore et nostrud
							cillum laborum. Labore reprehenderit labore mollit ad ea. Labore <b>qui</b>s fugiat
							nostrud aute deserunt sunt deserunt laborum elit eu aute esse aute. Irure elit
							<b>non</b>
							ipsum ali<b>qui</b>p mollit
							<b>non</b>
							id adipisicing sunt ad irure irure sit. <b>non</b> sint consectetur ex mollit elit sint.
						</p>
					</div>
					<div>
						<div>
							<Experience
								duration="SEPT 2023 - PRESENT"
								role="Software Engineering"
								company={{ title: 'Codeworks', website: 'https://codeworks.me/' }}
								level="Student"
								about="Deliver high-quality, robust production code for a diverse array of projects for clients including Harvard Business School, Everytown for Gun Safety, Pratt Institute, Koala Health, Vanderbilt University, The 19th News, and more. Provide leadership within engineering department through close collaboration, knowledge shares, and mentorship."
								skills={[
									'Javascript',
									'HTML',
									'CSS',
									'Node',
									'Express',
									'Koa',
									'GraphQL',
									'MongoDB',
									'Redis',
									'DevOps'
								]}
							/>
							<Experience
								duration="JULY 2023"
								role="Fontend Developer"
								company={{ title: 'Talon Outdoor', website: 'https://talonooh.com/en/' }}
								level="Internship"
								about="Deliver high-quality, robust production code for a diverse array of projects for clients including Harvard Business School, Everytown for Gun Safety, Pratt Institute, Koala Health, Vanderbilt University, The 19th News, and more. Provide leadership within engineering department through close collaboration, knowledge shares, and mentorship."
								skills={['Next.js', 'Jest', 'Sass', 'Javascript', 'Typescript']}
							/>
						</div>
					</div>{/if}
			</div>
		</div>
	</div>
</div>

<style lang="postcss">
	@import url('https://fonts.googleapis.com/css2?family=Inter:wght@100;200;300;400;500;600;700;800;900&display=swap');
	:global(html) {
		font-family: 'Inter', sans-serif;
		/* background: rgb(255, 0, 0); */
		background: rgb(2, 0, 36);
		background: linear-gradient(45deg, rgba(2, 0, 20, 1) 25%, rgba(40, 43, 48, 1) 100%);
	}
	p b {
		color: white;
	}
</style>
