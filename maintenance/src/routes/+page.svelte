<script lang="ts">
	import { Control, Field } from 'formsnap';
	import { superForm } from 'sveltekit-superforms';
	import type { PageData } from './$types';

	export let data: PageData;

	const form = superForm(data.requestForm, {
		resetForm: false
	});
	const { form: formData, enhance, message } = form;
</script>

<section class="grid h-full grid-rows-[auto_1fr] gap-2">
	<form method="post" use:enhance>
		<div class="grid grid-cols-[auto_1fr_auto] gap-3 items-center">
			<Field {form} name="authenticated">
				<Control let:attrs>
					<label for="authenticated" class="label flex flex-col">
						<span class="font-bold text-sm">Authenticated</span>
						<input
							{...attrs}
							class="input checkbox"
							type="checkbox"
							disabled={!data.user}
							bind:checked={$formData.authenticated}
						/>
					</label>
				</Control>
			</Field>
			<Field {form} name="action">
				<Control let:attrs>
					<select class="input select" {...attrs} bind:value={$formData.action}>
						{#each data.availableActions as action}
							<option value={action}>{action}</option>
						{/each}
					</select>
				</Control>
			</Field>
			<button type="submit" class="btn variant-filled-primary">Submit</button>
		</div>
	</form>
	<pre class="border border-surface-800 rounded-container-token">{JSON.stringify(
			typeof $message === 'string' ? JSON.parse($message) : $message,
			null,
			2
		)}</pre>
</section>
