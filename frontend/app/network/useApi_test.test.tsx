import { it, expect, describe, beforeEach } from 'vitest'
import { cleanup, render, screen } from '@testing-library/react'
import '@testing-library/jest-dom/vitest'
import { getMessageType } from './useApi'

describe('useApi', ()=>{

	beforeEach(()=> {
		cleanup();
	})

		it('should return "sucess" on 2** http status code', () => {
			expect(getMessageType(201)).toBe("sucess")
		})

		it('should return "error" on 4** http status code', () => {
			expect(getMessageType(401)).toBe("error")
		})

		it('should return "error" on 5** http status code', () => {
			expect(getMessageType(501)).toBe("error")
		})

		it('should return "info" by default', () => {
			expect(getMessageType(101)).toBe("info")
			expect(getMessageType(901)).toBe("info")
		})
	}
)
