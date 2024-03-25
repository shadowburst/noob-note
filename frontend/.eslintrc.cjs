const ERROR = 'error';
const WARN = 'warn';
const DISABLE = 'disable';

module.exports = {
    env: {
        browser: true,
        es2021: true,
    },
    extends: ['eslint:recommended', 'plugin:@typescript-eslint/recommended', 'plugin:vue/vue3-essential'],
    overrides: [
        {
            env: {
                node: true,
            },
            files: ['.eslintrc.{js,cjs}'],
            parserOptions: {
                sourceType: 'script',
            },
        },
    ],
    parserOptions: {
        ecmaVersion: 'latest',
        parser: '@typescript-eslint/parser',
        sourceType: 'module',
    },
    plugins: ['@typescript-eslint', 'vue'],
    rules: {
        camelcase: ERROR,
        'default-case': ERROR,
        'default-case-last': ERROR,
        eqeqeq: [WARN, 'smart'],
        indent: [ERROR, 4],
        'linebreak-style': [ERROR, 'unix'],
        'no-duplicate-imports': ERROR,
        'no-nested-ternary': ERROR,
        'no-var': ERROR,
        'prefer-arrow-callback': WARN,
        'prefer-const': WARN,
        'sort-imports': WARN,
        'sort-keys': WARN,
        'vue/multi-word-component-names': DISABLE,
    },
};
