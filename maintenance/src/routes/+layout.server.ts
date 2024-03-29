import { loginSchema } from '$lib/schemas/login-schema';
import { zod } from 'sveltekit-superforms/adapters';
import { superValidate } from 'sveltekit-superforms/server';
import type { LayoutServerLoad } from './$types';

export const load = (async ({ locals: { pb } }) => {
	const user = pb.authStore.model;
	const form = await superValidate(zod(loginSchema));
	return {
		user,
		form
	};
}) satisfies LayoutServerLoad;
