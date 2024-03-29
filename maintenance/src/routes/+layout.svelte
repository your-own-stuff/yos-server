<script lang="ts">
	import { arrow, autoUpdate, computePosition, flip, offset, shift } from '@floating-ui/dom';
	import { Toast, initializeStores, storePopup } from '@skeletonlabs/skeleton';
	import '../app.pcss';
	import type { LayoutData } from './$types';
	import Header from './header.svelte';

	export let data: LayoutData;

	storePopup.set({ computePosition, autoUpdate, offset, shift, flip, arrow });

	initializeStores();
</script>

<Toast position="br" />
<div class="grid h-full grid-rows-[auto_1fr] gap-2">
	<Header user={data.user?.username ?? null} data={data.form} />
	<main class="p-4">
		<slot />
	</main>
</div>
