<script>
	// @ts-nocheck
	import { showStats } from './stores/stats';
	import { wsConnection } from './stores/ws';
	import { clicks } from '../routes/stores/clicks';
	import FaGithub from 'svelte-icons/fa/FaGithub.svelte';
	import FaLinkedinIn from 'svelte-icons/fa/FaLinkedinIn.svelte';
	import FaGooglePlay from 'svelte-icons/fa/FaGooglePlay.svelte';
	import FaYoutube from 'svelte-icons/fa/FaYoutube.svelte';
	import LinkComponent from './link.svelte';
	import Experience from './experience.svelte';
	import Clickme from './clickme.svelte';
	import Map from './map.svelte';
	import { onMount } from 'svelte';
	import Bubbles from './bubbles.svelte';
	import { heatmap } from './stores/heatmap';
	import { selected } from './stores/selected';
	import Connections from './connections.svelte';
	import Projects from './projects.svelte';
	import HyperionImage from '$lib/assets/hyperion.png';
	import ErosImage from '$lib/assets/erosai.png';
	import { env } from '$env/dynamic/public';
	let loading = undefined;
	let more = false;

	function followMouse(event) {
		let elem = document.querySelector('#cursor');
		if (elem !== null) {
			let y = event.clientY - elem.offsetHeight / 2;
			let x = event.clientX - elem.offsetWidth / 2;
			elem.style.setProperty('left', x + 'px');
			elem.style.setProperty('top', y + 'px');
		}
	}

	onMount(async () => {
		loading = true;
		if (env.PUBLIC_DISABLE_ANIMATION != 'true') {
			let loading_id = setInterval(() => {
				loading = false;
				clearInterval(loading_id);
			}, 2500);
		} else {
			loading = false;
		}

		function connect() {
			console.log('connecting...');
			const ws = new WebSocket(`${env.PUBLIC_WS_ENDPOINT}/ws/actions?v=1`);
			ws.addEventListener('open', () => {
				console.log('connected');
				$clicks['disabled'] = false;
				$wsConnection = ws;
			});
			ws.addEventListener('message', (message) => {
				// Parse the incoming message here
				if (message.data !== '') {
					const data = JSON.parse(message.data);
					console.log('recieved', data);
					if (data.type === 'clicks') {
						$clicks['clicks'] = data.value;
					} else if (data.type === 'stats') {
						$heatmap = JSON.parse(data.value);
					} else if (data.type === 'error') {
						$clicks['disabled'] = true;
						setTimeout(() => {
							$clicks['disabled'] = false;
						}, 5000);
					}
				} else {
					console.log('recieved', message);
				}
			});
			ws.addEventListener('close', () => {
				console.log('closed connection');
				$clicks['disabled'] = true;
				setTimeout(function () {
					connect();
				}, 1000);
			});
		}
		connect();
	});
</script>

<svelte:head>
	<title>Alexander Heather</title>
	<meta charset="utf-8" />
	<meta name="author" content="Alexander Heather" />
	<meta
		name="description"
		content="My personal portfolio for showing off my projects and experiences. And also demonstrating my web development skills"
	/>
</svelte:head>

{#if loading === undefined || loading}
	{#if loading}
		<div class="m-auto items-center flex h-screen min-h-full w-screen mx-auto relative">
			<div class=" text-white">
				<svg
					class="path"
					stroke="#fff"
					stroke-width="2px"
					width="274"
					height="188"
					viewBox="0 0 274 188"
					fill="#fff"
					fill-opacity="0%"
					xmlns="http://www.w3.org/2000/svg"
				>
					<path
						d="M42.5 187.5H1L10 165H34.5L118 0.5H151.5L167 37L138 93H118L146.5 37L138 22.5H125.5L42.5 187.5Z"
						fill="white"
						stroke="white"
					/>
					<path
						d="M102 186L174 48.5L236 165H262L272.5 186H223.5L174 93L125 186H102Z"
						fill="white"
						stroke="white"
					/>
				</svg>
			</div>
		</div>
	{/if}
{:else}
	<div
		id="cursor"
		class="bg-gradient-to-r fixed from-s-purple to-s-blue w-32 h-32 bg-blend-multiply blur-3xl opacity-60 -z-50 hidden md:block"
	/>
	<!-- svelte-ignore a11y-no-static-element-interactions -->
	<div id="main" class="relative" on:mousemove={followMouse}>
		<Bubbles left="5%" top="5%" classExtra="hidden md:block" />
		<Bubbles left="93%" top="85%" />
		<div class="blur-3xl -top-5 -left-5 w-16 h-16 md:w-32 md:h-32 fixed bg-s-purple" />

		<div class="min-h-screen w-screen md:pt-20 md:px-32 text-s-light-gray p-6 md:p-0">
			<div class="md:grid md:grid-cols-2 gap-x-12">
				<div class="md:fixed md:w-1/2 pb-12 md:pb-0">
					<div class="">
						<h1 class="text-5xl lg:text-7xl text-white font-bold">ALEXANDER HEATHER</h1>
						<h2 class="text-3xl lg:text-5xl text-white font-medium text-opacity-90 pb-6 pt-2">
							WEB DEVELOPER
						</h2>
						<h3 class="text-2xl lg:text-4xl font-medium md:w-2/3">
							I excel in crafting high-performance websites that seamlessly integrate stunning
							visual elements
						</h3>
					</div>
					<!-- <div class="pt-16 w-1/2">
						<LinkComponent linkName="about" location="0%" />
						<LinkComponent linkName="projects" location="50%" />
						<LinkComponent linkName="experiences" location="100%" />
					</div> -->
					<div class="pt-8 md:pt-16">
						<Clickme />
					</div>
					<div class="flex justify-between md:justify-start md:gap-x-4 lg:gap-x-8 pt-12">
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
				<div class="z-10" id="scrollable">
					<!-- <Map /> -->
					{#if $showStats}
						<Map />
					{:else}
						<div class="pb-6 md:pb-0">
							<h1 class="opacity-0" id="about">About</h1>
							<p class="text-xl md:text-sm lg:text-xl font-normal pb-6">
								In <b>2020</b>, I embarked on a journey to become a web developer, beginning my
								journey with creating sites utilising <b>pug.js</b>. This was an enlightening first
								step and led to my exploration of different programming languages and frameworks
								such as
								<b>Python</b> and <b>React.js</b> which I employed to develop utility programs and
								improve my everyday tasks. Aswell as make websites<br /><br />By the time
								<b>2022</b> came around, I had grown more confident about programming skills and
								expanded my knowledge base to include <b>golang</b> and <b>dart</b>. Additionally,
								concepts such as
								<span class={'md:block ' + (!more ? 'hidden' : 'block')}
									><b>kubernetes</b> and <b>devops</b> also found a place in my field of expertise.<br
									/><br />May of
									<b>2023</b> marked a significant milestone in my coding career as I launched my
									first production application. This application was programmed in <b>Dart</b> using
									the <b>Flutter</b>
									framework, and its backend was supported by a fusion of <b>golang</b> and
									<b>python</b>
									arranged in a microservices architecture.<br /><br />I further ventured into
									sophisticated concepts such as <b>gRPC</b> for enhanced communication between services.
									This quest for knowledge spilled over into self-hosting services like object storage,
									and a GitLab instance and elastic search for swiftly exploring a substantial volume
									of log files.</span
								>
							</p>
							<button
								on:click={(target) => {
									if (more) {
										let el = document.querySelector('#about');
										// if (!el) return;
										el.scrollIntoView({
											behavior: 'smooth'
										});
									}

									more = !more;
								}}
								class="mx-auto md:hidden md:pb-0 flex bg-s-blue bg-opacity-10 text-s-blue px-12 text-md py-2 rounded-full font-semibold"
								>{more ? 'Less' : 'More'}</button
							>
						</div>
						<div class="pb-12">
							<div class="pb-12 grid grid-cols-1 gap-y-12">
								<Projects
									title="Eros AI"
									duration="MAY 2023 - PRESENT"
									about="Eros AI started out as a test to gauge my abilities. I ended up liking where it was going and decided to make it into an actual app that is publicly available. Since it's release in may it has gained over 5,000 downloads and continues to increase daily. It has also generated over 57,000 AI images overall."
									img={ErosImage}
									connections={[
										{
											href: 'https://play.google.com/store/apps/details?id=net.atlasservers.eros',
											icon: FaGooglePlay
										},
										{
											href: 'https://www.youtube.com/@eros_ai/shorts',
											icon: FaYoutube
										}
									]}
									href="https://github.com/SneakyPanda64/Hyperion"
									skills={[
										'Golang',
										'Python',
										'Dart',
										'Flutter',
										'Kubernetes',
										'MinIO',
										'gRPC',
										'NextJs',
										'Microservices',
										'Stable Diffusion',
										'PyTorch',
										'Tensorflow'
									]}
								/>
								<Projects
									title="Hyperion"
									duration="OCT - NOV 2022"
									about="Hyperion was a project developed to automatically create YouTube videos that would then get uploaded at set intervals daily. I was the sole developer of the whole project. You can checkout the videos that were created on my YouTube channel. For details about how Hyperion works checkout the README.md"
									img={HyperionImage}
									connections={[
										{
											href: 'https://github.com/SneakyPanda64/Hyperion',
											icon: FaGithub
										},
										{
											href: 'https://www.youtube.com/@atlas_tv',
											icon: FaYoutube
										}
									]}
									href="https://github.com/SneakyPanda64/Hyperion"
									skills={['Python', 'Deep Learning', 'GPT3', 'FFmpeg', 'Docker', 'FTP']}
								/>
							</div>
							<div>
								<!-- <Experience
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
								/> -->
								<Projects
									duration="JULY 2023"
									title="Fontend Developer • Talon Outdoor"
									subtitle="Internship"
									href="https://talonooh.com/"
									about="I had the opportunity to work in a professional setting alongside other web developers, where my main focus was on developing a NextJS web application. During my time at Talon, I was able to contribute several new features and resolve numerous bugs. Working alongside a highly skilled team of individuals, I had the privilege of learning from their expertise and gaining valuable insights."
									skills={['Next.js', 'Jest', 'Sass', 'Javascript', 'Typescript']}
								/>
							</div>
						</div>{/if}
				</div>
			</div>
		</div>
	</div>
{/if}

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
		font-weight: normal;
	}
	.path {
		stroke-dasharray: 1000;
		stroke-dashoffset: 1000;
		animation: dash 3s linear forwards;
		position: absolute;
		left: 50%;
		top: 50%;
		transform: translate(-50%, -50%);
	}

	@keyframes dash {
		100% {
			stroke-dashoffset: 0;
		}
	}
	@keyframes opacity {
		50% {
			fill-opacity: 0%;
		}
		100% {
			fill-opacity: 100%;
		}
	}
</style>
