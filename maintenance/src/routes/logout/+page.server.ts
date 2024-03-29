import { type Actions } from '@sveltejs/kit';

export const actions: Actions = {
	default: async ({ locals: { pb } }) => {
		pb.authStore.clear();
	}
};
