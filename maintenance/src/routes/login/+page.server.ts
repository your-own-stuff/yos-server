import { loginSchema } from '$lib/schemas/login-schema';
import type { Actions } from '@sveltejs/kit';
import { zod } from 'sveltekit-superforms/adapters';
import { setError, superValidate } from 'sveltekit-superforms/server';

export const actions: Actions = {
	default: async ({ request, locals: { pb } }) => {
		const form = await superValidate(request, zod(loginSchema));

		if (!form.valid) return form;

		try {
			await pb.collection('users').authWithPassword(form.data.email, form.data.password);
		} catch (e) {
			console.log(e);

			return setError(form, 'password', 'Invalid email or password');
		}
	}
};
