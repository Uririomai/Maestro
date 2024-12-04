import React from 'react'

import styles from './button.module.scss'

const Button = props => {
	const styleElements = {
		background: props.colorBack,
		color: props.colorText,
		height: props.height,
		width: props.width,
		
	}

	if (props.imageSrc) {
		return (
			<>
				<button
					className={styles.button}
					style={styleElements}
					type='button'
					onClick={props.onClick}
				>
					<img
						className={styles.button__img}
						src={props.imageSrc}
						alt='img-button'
					/>

					<p className={styles.button__text}>{props.buttonText}</p>
				</button>
			</>
		)
	} else {
		return (
			<>
				<button
					className={styles.button}
					style={styleElements}
					type='button'
					onClick={props.onClick}
				>
					<p className={styles.button__text}>{props.buttonText}</p>
				</button>
			</>
		)
	}
}

export default Button
