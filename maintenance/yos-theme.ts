import type { CustomThemeConfig } from '@skeletonlabs/tw-plugin';

export const yosTheme: CustomThemeConfig = {
	name: 'yos-theme',
	properties: {
		// =~= Theme Properties =~=
		'--theme-font-family-base': `system-ui, SFMono-Regular, Menlo, Monaco, Consolas, 'Liberation Mono', 'Courier New', monospace`,
		'--theme-font-family-heading': `system-ui`,
		'--theme-font-color-base': '0 0 0',
		'--theme-font-color-dark': '255 255 255',
		'--theme-rounded-base': '12px',
		'--theme-rounded-container': '8px',
		'--theme-border-base': '1px',
		// =~= Theme On-X Colors =~=
		'--on-primary': '255 255 255',
		'--on-secondary': '255 255 255',
		'--on-tertiary': '0 0 0',
		'--on-success': '0 0 0',
		'--on-warning': '0 0 0',
		'--on-error': '255 255 255',
		'--on-surface': '255 255 255',
		// =~= Theme Colors  =~=
		// primary | #6100DE
		'--color-primary-50': '231 217 250', // #e7d9fa
		'--color-primary-100': '223 204 248', // #dfccf8
		'--color-primary-200': '216 191 247', // #d8bff7
		'--color-primary-300': '192 153 242', // #c099f2
		'--color-primary-400': '144 77 232', // #904de8
		'--color-primary-500': '97 0 222', // #6100DE
		'--color-primary-600': '87 0 200', // #5700c8
		'--color-primary-700': '73 0 167', // #4900a7
		'--color-primary-800': '58 0 133', // #3a0085
		'--color-primary-900': '48 0 109', // #30006d
		// secondary | #335462
		'--color-secondary-50': '224 229 231', // #e0e5e7
		'--color-secondary-100': '214 221 224', // #d6dde0
		'--color-secondary-200': '204 212 216', // #ccd4d8
		'--color-secondary-300': '173 187 192', // #adbbc0
		'--color-secondary-400': '112 135 145', // #708791
		'--color-secondary-500': '51 84 98', // #335462
		'--color-secondary-600': '46 76 88', // #2e4c58
		'--color-secondary-700': '38 63 74', // #263f4a
		'--color-secondary-800': '31 50 59', // #1f323b
		'--color-secondary-900': '25 41 48', // #192930
		// tertiary | #D93D22
		'--color-tertiary-50': '249 226 222', // #f9e2de
		'--color-tertiary-100': '247 216 211', // #f7d8d3
		'--color-tertiary-200': '246 207 200', // #f6cfc8
		'--color-tertiary-300': '240 177 167', // #f0b1a7
		'--color-tertiary-400': '228 119 100', // #e47764
		'--color-tertiary-500': '217 61 34', // #D93D22
		'--color-tertiary-600': '195 55 31', // #c3371f
		'--color-tertiary-700': '163 46 26', // #a32e1a
		'--color-tertiary-800': '130 37 20', // #822514
		'--color-tertiary-900': '106 30 17', // #6a1e11
		// success | #84cc16
		'--color-success-50': '237 247 220', // #edf7dc
		'--color-success-100': '230 245 208', // #e6f5d0
		'--color-success-200': '224 242 197', // #e0f2c5
		'--color-success-300': '206 235 162', // #ceeba2
		'--color-success-400': '169 219 92', // #a9db5c
		'--color-success-500': '132 204 22', // #84cc16
		'--color-success-600': '119 184 20', // #77b814
		'--color-success-700': '99 153 17', // #639911
		'--color-success-800': '79 122 13', // #4f7a0d
		'--color-success-900': '65 100 11', // #41640b
		// warning | #EAB308
		'--color-warning-50': '252 244 218', // #fcf4da
		'--color-warning-100': '251 240 206', // #fbf0ce
		'--color-warning-200': '250 236 193', // #faecc1
		'--color-warning-300': '247 225 156', // #f7e19c
		'--color-warning-400': '240 202 82', // #f0ca52
		'--color-warning-500': '234 179 8', // #EAB308
		'--color-warning-600': '211 161 7', // #d3a107
		'--color-warning-700': '176 134 6', // #b08606
		'--color-warning-800': '140 107 5', // #8c6b05
		'--color-warning-900': '115 88 4', // #735804
		// error | #e12323
		'--color-error-50': '251 222 222', // #fbdede
		'--color-error-100': '249 211 211', // #f9d3d3
		'--color-error-200': '248 200 200', // #f8c8c8
		'--color-error-300': '243 167 167', // #f3a7a7
		'--color-error-400': '234 101 101', // #ea6565
		'--color-error-500': '225 35 35', // #e12323
		'--color-error-600': '203 32 32', // #cb2020
		'--color-error-700': '169 26 26', // #a91a1a
		'--color-error-800': '135 21 21', // #871515
		'--color-error-900': '110 17 17', // #6e1111
		// surface | #424243
		'--color-surface-50': '227 227 227', // #e3e3e3
		'--color-surface-100': '217 217 217', // #d9d9d9
		'--color-surface-200': '208 208 208', // #d0d0d0
		'--color-surface-300': '179 179 180', // #b3b3b4
		'--color-surface-400': '123 123 123', // #7b7b7b
		'--color-surface-500': '66 66 67', // #424243
		'--color-surface-600': '59 59 60', // #3b3b3c
		'--color-surface-700': '50 50 50', // #323232
		'--color-surface-800': '40 40 40', // #282828
		'--color-surface-900': '32 32 33' // #202021
	}
};
