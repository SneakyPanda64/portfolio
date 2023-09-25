<script>
	import { selected } from './stores/selected';
	export let linkName = '';
	const onClick = (e) => {
		console.log($selected);
		if ($selected === linkName) {
			$selected = undefined;
		}
		$selected = linkName;
	};
</script>

<a href={`#${linkName}`} on:click={onClick} class="group cursor-pointer">
	<div
		class={'flex transition-all duration-300 links ' +
			($selected === linkName ? 'selected' : 'unselected')}
	>
		<div
			class="h-[0.1rem] my-auto bg-s-blue w-full bg-opacity-50 group-hover:bg-opacity-100 mr-4"
		/>
		<h1 class="text-lg font-semibold group-hover:text-white">{linkName.toUpperCase()}</h1>
	</div>
</a>

<style>
	.links {
		width: 50%;
	}
	.selected {
		animation: extend 500ms forwards;
	}
	.unselected {
		animation: shrink 200ms forwards;
	}
	@keyframes extend {
		0% {
			width: 50%;
		}
		100% {
			width: 75%;
			color: white;
		}
	}
	@keyframes shrink {
		0% {
			width: 75%;
		}
		100% {
			width: 50%;
			color: #6f747a;
		}
	}
</style>
