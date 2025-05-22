import { expect, it } from "vitest";
import { render } from "@testing-library/react";
import '@testing-library/jest-dom';
import Home from "./page";
import { describe } from "node:test";

describe("Home Page", () => {
	it('should render Home', () => {
		const { getByText } = render(<Home />);
	expect(getByText('Hello')).toBeInTheDocument();
  });
})
