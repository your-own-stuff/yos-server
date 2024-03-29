import { createPb } from '$lib/create-pb';
import type { Actions } from '@sveltejs/kit';
import { message, superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import { z } from 'zod';
import type { PageServerLoad } from './$types';

const availableActions = ['rebuild-index', 'something-else'] as const;

const pocketBaseRequest = z.object({
	action: z.enum(availableActions),
	authenticated: z.boolean()
});

export const load = (async ({ locals: { pb } }) => {
	const requestForm = await superValidate(
		{ authenticated: pb.authStore.model ? true : false },
		zod(pocketBaseRequest)
	);
	return { requestForm, availableActions };
}) satisfies PageServerLoad;

export const actions: Actions = {
	default: async ({ request, locals: { pb } }) => {
		const form = await superValidate(request, zod(pocketBaseRequest));
		const sendWith = form.data.authenticated ? pb : createPb();

		try {
			const retVal = await sendWith.send(`/${form.data.action}`, {});
			return message(form, retVal);
		} catch (e) {
			return message(form, JSON.stringify(e));
		}
	}
};
