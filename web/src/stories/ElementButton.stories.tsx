import React from "react";
// also exported from '@storybook/react' if you can deal with breaking changes in 6.1
import { Story, Meta } from "@storybook/react/types-6-0";

import { Element, ElementProps } from "../components/Element";

export default {
  title: "Example/ElementButton",
  component: Element,
  argTypes: {
    backgroundColor: { control: "color" },
  },
} as Meta;

const Template: Story<ElementProps> = (args) => <Element {...args} />;

export const Large = Template.bind({});
Large.args = {
  children: "My Element Button",
};

export const Medium = Template.bind({});
Medium.args = {
  children: "My Element Button",
  size: "medium",
};

export const Small = Template.bind({});
Small.args = {
  children: "My Element Button",
  size: "small",
};

export const Active = Template.bind({});
Active.args = {
  children: "Active",
  size: "small",
  active: true,
};

export const Passive = Template.bind({});
Passive.args = {
  children: "Passive",
  size: "small",
};
