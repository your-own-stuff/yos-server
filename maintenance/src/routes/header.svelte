<script lang="ts">
	import { enhance } from '$app/forms';
	import { type LoginSchema } from '$lib/schemas/login-schema';
	import { type Infer, type SuperValidated } from 'sveltekit-superforms';
	import LoginForm from './login-form.svelte';

	export let user: string | null;
	export let data: SuperValidated<Infer<LoginSchema>>;
</script>

<header class="border-b border-b-surface-800">
	<div class="flex justify-between items-center p-2">
		<span class="h3">Maintenance</span>
		{#if user}
			<div class="flex gap-2 items-center">
				Authenticated as {user}
				<form action="/logout" method="post" use:enhance>
					<button class="btn btn-sm variant-filled-error">Logout</button>
				</form>
			</div>
		{:else}
			<LoginForm {data} />
		{/if}
	</div>
</header>
