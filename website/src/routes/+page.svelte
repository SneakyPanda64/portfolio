<script>
	// @ts-nocheck
	import { showStats } from './stores/stats';
	import LinkComponent from './link.svelte';
	import Experience from './experience.svelte';
	import Clickme from './clickme.svelte';
	import Map from './map.svelte';
	// @ts-ignore
	function followMouse(event) {
		let elem = document.querySelector('#cursor');
		if (elem !== null) {
			let y = event.pageY - elem.offsetHeight / 2;
			let x = event.pageX - elem.offsetWidth / 2;
			// let x = (event.clientX - elem.offsetWidth) / 2;
			// @ts-ignore
			elem.style.setProperty('position', 'absolute');
			// @ts-ignore
			elem.style.setProperty('left', x + 'px');
			// @ts-ignore
			elem.style.setProperty('top', y + 'px');
		}
	}
</script>

<!-- svelte-ignore a11y-no-static-element-interactions -->
<div class="relative" on:mousemove={followMouse}>
	<div
		id="cursor"
		class="bg-gradient-to-r from-s-purple to-s-blue w-32 h-32 absolute blur-3xl opacity-60 -z-50"
	/>
	<div class="min-h-screen w-screen pt-24 px-32 text-s-light-gray">
		<div class="grid grid-cols-2 gap-x-12">
			<div class="fixed w-1/2">
				<div>
					<h1 class="text-5xl text-white font-bold">ALEXANDER HEATHER</h1>
					<h2 class="text-4xl text-white font-medium pb-6">WEB DEVELOPER</h2>
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
			</div>
			<div />
			<div>
				<!-- <Map /> -->
				{#if $showStats}
					<Map />
				{:else}
					<div>
						<p class="text-xl font-normal">
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
							Veniam consequat nostrud officia nostrud reprehenderit ipsum nisi proident aute nisi Lorem.
							Pariatur nulla adipisicing id laboris nisi ullamco magna excepteur mollit ex deserunt consequat
							labore eu. Voluptate magna eu veniam ex nulla excepteur enim sit. Anim elit commodo reprehenderit
							aliqua id veniam elit pariatur minim fugiat.
						</p>
					</div>
					<div>
						<div>
							<Experience />
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
		background: linear-gradient(
			150deg,
			rgba(2, 0, 36, 1) 0%,
			rgba(40, 43, 48, 1) 90%,
			rgba(61, 71, 106, 0.979) 100%
		);
	}
	p b {
		color: white;
	}
</style>
