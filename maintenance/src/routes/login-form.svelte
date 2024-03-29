<script lang="ts">
	import { loginSchema, type LoginSchema } from '$lib/schemas/login-schema';
	import { getToastStore } from '@skeletonlabs/skeleton';
	import { Control, Field } from 'formsnap';
	import { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import { zodClient } from 'sveltekit-superforms/adapters';
	const toast = getToastStore();

	export let data: SuperValidated<Infer<LoginSchema>>;
	const form = superForm(data, {
		validators: zodClient(loginSchema),
		onUpdate: ({ form }) => {
			if (!form.valid) {
				toast.trigger({ message: 'Invalid login' });
			}
		}
	});

	const { form: formData, enhance } = form;
</script>

<form method="post" action="/login" use:enhance class="flex items-end gap-2">
	<Field {form} name="email">
		<Control let:attrs>
			<input placeholder="user" class="input text-xs" {...attrs} bind:value={$formData.email} />
		</Control>
	</Field>
	<Field {form} name="password">
		<Control let:attrs>
			<input
				placeholder="password"
				class="input text-xs"
				{...attrs}
				type="password"
				bind:value={$formData.password}
			/>
		</Control>
	</Field>
	<button class="variant-filled-primary btn text-xs">Login</button>
</form>
