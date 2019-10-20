// DONE
describe('Check the load of all pages', function () {
  it('Home page load', function () {
    cy.visit('')
    cy.contains('Self-hosted web application for weekly reporting')
  })

  it('Help page loads', function() {
    cy.visit('help')
    cy.contains('Quick Start')
  })

  it('Documentation page load', function () {
    cy.visit('documentation')
    cy.contains('Requirements')
  })

  it('Administration page load', function () {
    cy.visit('administration')
    cy.contains('Manage Consultants')
  })

  it('Holiday page load', function () {
    cy.visit('holidays')
    cy.contains('State Holidays')
  })

  it('Overview page load', function () {
    cy.visit('overview')
    cy.contains('Reported Time')
    cy.contains('Reported Projects')
  })

  it('Report page load', function () {
    cy.visit('report')
    cy.contains('Weekly')
  })
})
