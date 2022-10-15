import { component$, useStylesScoped$ } from "@builder.io/qwik";
import styles from "./header.css?inline";

export default component$(() => {
	useStylesScoped$(styles);

	return (
		<header>
			<h3>Dion</h3>
		</header>
	);
});
