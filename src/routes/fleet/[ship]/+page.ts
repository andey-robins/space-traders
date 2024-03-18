import { error } from '@sveltejs/kit';
import type { PageData, PageLoad } from './$types';

export const load: PageLoad = ({ params }: { params: PageData }) => {
	if (params.ship === 'test') {
		return {
			callsign: 'TEST_SHIP'
		};
	}

	error(404, 'Not found');
};
