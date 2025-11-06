// @ts-check
import eslint from '@eslint/js'
import nextPlugin from '@next/eslint-plugin-next'
import tanstackPlugin from '@tanstack/eslint-plugin-query'
import prettierConfig from 'eslint-config-prettier'
import * as importPlugin from 'eslint-plugin-import'
import hooksPlugin from 'eslint-plugin-react-hooks'
import reactPluginConfigJSXRuntime from 'eslint-plugin-react/configs/jsx-runtime.js'
import reactPluginConfigRecommended from 'eslint-plugin-react/configs/recommended.js'
import storybookPlugin from 'eslint-plugin-storybook'
import unicornPlugin from 'eslint-plugin-unicorn'
import unusedImportsPlugin from 'eslint-plugin-unused-imports'
import tsEslint from 'typescript-eslint'

export default [
  {
    ignores: [
      'tailwind.config.ts',
      'vitest.config.ts',
      '*.test.ts',
      'next.config.js',
      '.next/**/*',
      'postcss.config.js'
    ]
  },
  eslint.configs.recommended,
  ...tsEslint.configs.recommended,
  ...tanstackPlugin.configs['flat/recommended'],
  ...storybookPlugin.configs['flat/recommended'],
  reactPluginConfigRecommended,
  reactPluginConfigJSXRuntime,
  prettierConfig,
  {
    languageOptions: {
      parser: tsEslint.parser,
      parserOptions: {
        projectService: true,
        project: './tsconfig.json'
      }
    },
    settings: {
      react: {
        version: 'detect'
      }
    },
    files: ['**/*.ts', '**/*.tsx', '**/*.js'],
    plugins: {
      'react-hooks': hooksPlugin,
      '@next/next': nextPlugin,
      'unused-imports': unusedImportsPlugin,
      unicorn: unicornPlugin,
      reactPluginConfigRecommended,
      import: importPlugin
    },
    rules: {
      ...hooksPlugin.configs.recommended.rules,
      ...nextPlugin.configs.recommended.rules,
      ...nextPlugin.configs['core-web-vitals'].rules,
      'no-fallthrough': 'error',
      'no-console': ['error', { allow: ['warn', 'error', 'info'] }],
      'no-debugger': 'error',
      // イベントハンドラの Props 名は onXxx、イベントハンドラ関数名は handleXxx とする
      'react/jsx-handler-names': [
        'error',
        {
          // ローカル変数のイベントハンドラもチェック
          checkLocalVariables: true,
          // インライン関数のイベントハンドラをチェック
          checkInlineFunction: true
        }
      ],
      'unicorn/filename-case': [
        'error',
        {
          cases: {
            kebabCase: true,
            camelCase: false,
            snakeCase: false,
            pascalCase: false
          }
        }
      ],
      '@typescript-eslint/consistent-type-imports': [
        'error',
        {
          // import type を強制する
          prefer: 'type-imports',
          // --fix 実行時にインラインスタイル（import { type ... }）を適用する
          fixStyle: 'inline-type-imports'
        }
      ],
      // 型定義は interface を強制する
      '@typescript-eslint/consistent-type-definitions': 'error',
      // インラインスタイルの import type を強制する
      'import/consistent-type-specifier-style': ['error', 'prefer-inline'],
      // 同じパスからの import を禁止する
      'import/no-duplicates': [
        'error',
        {
          // --fix 実行時にインラインスタイルを適用する
          'prefer-inline': true
        }
      ],
      // import の順序を 依存パッケージ、絶杯パス、相対パスの順番に強制する
      'import/order': [
        'error',
        {
          groups: [['builtin', 'external'], 'internal', ['parent', 'sibling', 'index']],
          'newlines-between': 'always',
          pathGroups: [
            {
              pattern: '@/**',
              group: 'internal',
              position: 'before'
            }
          ]
        }
      ]
    }
  },
  {
    files: ['**/*.ts'],
    rules: {
      // ts ファイルの戻り値の型を強制する
      '@typescript-eslint/explicit-function-return-type': 'error'
    }
  }
]
