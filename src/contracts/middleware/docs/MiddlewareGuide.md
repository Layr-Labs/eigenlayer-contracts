# Introduction
In designing EigenLayer, we aspired to make minimal assumptions about the structure of middlewares built on top of it. If you are getting started looking at the EigenLayer codebase, the Slasher contract contains most of the logic that actually mediates the interactions between EigenLayer and middlewares. Additionally, there is a general-purpose /middleware/ folder, which contains code that can be extended, used directly, or consulted as a reference in building middleware on top of EigenLayer.


# Assumptions EigenLayer Contracts Make About Middlewares
