// TODO
describe('Administration page is functional', function() {
  it('Information logs can be shown', function () {
      cy.visit('/administration')
    cy.get('div').contains('View Application Logs').click()
    cy.get('div').contains('Select log level').click()
    cy.get('div').contains('Information').click()
    cy.get('div[class="absolute full-width"]').contains('[INFO]')
  })

  it('Managed data table contains correct header and at least 1 record per each entity', function () {
    cy.visit('/administration')
    cy.get('div').contains('Managed Data Statistics').click()
    cy.get('table').should('be.visible')
    cy.get('table').get('th').contains('# of total records')
    cy.get('tbody').eq(2).children().should('have.length', 4)
    // check reported records
    cy.get('tbody').eq(2).children().eq(0).children().eq(1).invoke('text').then((text => {
      expect(text).to.not.eq('0')
    }))
    // check projects
    cy.get('tbody').eq(2).children().eq(1).children().eq(1).invoke('text').then((text => {
      expect(text).to.not.eq('0')
    }))
    // check rates
    cy.get('tbody').eq(2).children().eq(2).children().eq(1).invoke('text').then((text => {
      expect(text).to.not.eq('0')
    }))
    // check consultants
    cy.get('tbody').eq(2).children().eq(3).children().eq(1).invoke('text').then((text => {
      expect(text).to.not.eq('0')
    }))
  })

  // it('Data can be downloaded', function () {
  //   cy.visit('/administration')
  //   cy.get('div').contains('Backup & Restore').click()
  //   cy.get('button').contains('Download').click({force:true})
  // })

  // it('Create new consultant', function() {
  //   cy.visit('administration')

  //   cy.get('[data-cy="consultants_projects"]')
  //     .click()

  //   cy.get('[data-cy="consultant_name"]')
  //     .type('Stanislav')

  //   cy.get('[data-cy="consultant_create"]')
  //     .click()
  // })

  // it('Hide consultant', function () {
  //   cy.visit('/administration')
  // })

  // it('Show consultant', function () {
  //   cy.visit('/administration')
  // })

  // it('Delete consultant', function () {
  //   cy.visit('/administration')
  // })
})

// Administration of Project
// Create
// Hide
// Show
// Delete

// managed data statistics
// each entity count > 1

// Download data
// Upload downloaded data

// Show Info logs

