import { PUBLIC_YOS_URL } from '$env/static/public';
import PocketBase from 'pocketbase';
import type { YosPocketBase } from './types/yos-pocket-base';

export function createPb() {
	return new PocketBase(PUBLIC_YOS_URL) as YosPocketBase;
}
