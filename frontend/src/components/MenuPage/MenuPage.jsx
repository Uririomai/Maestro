import React from 'react'
import Sidebar from '../Sidebar/Sidebar'

import cabinetStyles from '../Cabinet/cabinet.module.scss'
import styles from './menuPage.module.scss'

const MenuPage = () => {
  return (
    <>
			<section className={cabinetStyles.personalCabinet + ` container`}> 
                <Sidebar />
                
            </section>
		</>
  )
}

export default MenuPage