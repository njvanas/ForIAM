import { describe, it, expect } from 'vitest'
import { hello } from './index'

describe('hello function', () => {
  it('greets properly', () => {
    expect(hello('world')).toBe('Hello, world!')
  })
})
