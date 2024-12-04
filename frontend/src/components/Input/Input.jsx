import React from 'react'

import styles from './input.module.scss'

export const Input = props => {
	return (
		<>
			<label className={styles.label} htmlFor={props.id}>
				{props.label}
				<input
					className={styles.label__input}
					type={props.type}
					name={props.id}
					id={props.id}
					disabled={props.isDisable}
					placeholder={props.placeholder}
				/>{' '}
			</label>
		</>
	)
}
export const TextArea = props => {
	return (
		<>
			<label className={styles.label} htmlFor={props.id}>
				{props.label}
				<textarea
					className={styles.label__textArea}
					type={props.type}
					name={props.id}
					id={props.id}
					disabled={props.isDisable}
					placeholder={props.placeholder}
					cols='40'
					rows='3'
				></textarea>
			</label>
		</>
	)
}
export const SelectIn = props => {
	return (
		<>
			<label className={styles.label} htmlFor={props.id}>
				{props.label}
				<select
					className={styles.label__input}
					type={props.type}
					name={props.id}
					id={props.id}
					disabled={props.isDisable}
					placeholder={props.placeholder}
					cols='40'
					rows='3'
				></select>
			</label>
		</>
	)
}
export const FileIn = props => {
	return (
		<>
			<label className={styles.label} htmlFor={props.id}>
				{props.label}
				<input
					className={styles.label__file}
					type={'file'}
					name={props.id}
					id={props.id}
					disabled={props.isDisable}
					placeholder={props.placeholder}
				></input>
			</label>
		</>
	)
}

export default Input
