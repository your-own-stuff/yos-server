// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
import type { YosPocketBase } from '$lib/types/yos-pocket-base';

declare global {
	namespace App {
		// interface Error {}
		interface Locals {
			pb: YosPocketBase;
		}
		// interface PageData {}
		// interface PageState {}
		// interface Platform {}
	}
}

export {};
