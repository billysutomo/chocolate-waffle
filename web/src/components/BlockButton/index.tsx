import React from 'react';

export interface BlockButtonProps {
    /**
     * Is this the principal call to action on the page?
     */
    // primary?: boolean;
    // /**
    //  * What background color to use
    //  */
    // backgroundColor?: string;
    // /**
    //  * How large should the button be?
    //  */
    // size?: 'small' | 'medium' | 'large';
    /**
     * Button contents
     */
    label: string;
    /**
     * Optional click handler
     */
    // onClick?: () => void;
  }

export const BlockButton: React.FC<BlockButtonProps> = ({
    label = "default",
    ...props
}) => {
    return (
    <span>{label}</span>
    )
}